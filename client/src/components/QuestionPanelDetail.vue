<template>
  <div class="question-panel-detail">
    <v-card
      class="mx-auto"
    >
      <div v-if="isReady">
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="display-1">
              {{ question.title }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider class="mx-4"></v-divider>
        <v-row>
          <v-col md="12"> 
            <v-card-text>
              <div class="mavon-editor">
                <mavon-editor 
                  v-model="question.body"
                  defaultOpen="preview"
                  :toolbarsFlag="fa"
                  :subfield="fa"
                  :boxShadow="fa"
                />
              </div>
            </v-card-text>
            <!--v-btn
              class="tag"
              color="blue-grey lighten-4"
              x-small
            >
              tag (仮)
            </v-btn>
            <v-btn
              class="tag"
              color="blue-grey lighten-4"
              x-small
            >
              tag (仮)
            </v-btn>
            <v-btn
              class="tag"
              color="blue-grey lighten-4"
              x-small
            >
              tag (仮)
            </v-btn-->
          </v-col>
        </v-row>
        <!--v-chip
          class="ma-2"
          color="blue"
          text-color="white"
        >
          <v-avatar
            left
            class="blue darken-4"
          >
            {{ question.answerCount }}
          </v-avatar>
          回答数
        </v-chip>
        <v-chip
          class="ma-2"
          color="teal"
          text-color="white"
        >
          <v-avatar
            left
            class="teal darken-4"
          >
            {{ question.browseCount }}
          </v-avatar>
          閲覧数
        </v-chip>
        <v-chip
          class="ma-2"
          color="pink"
          text-color="white"
        >
          <v-avatar
            left
            class="pink darken-4"
          >
            {{ question.favoriteCount }}
          </v-avatar>
          いいね
        </v-chip-->
        <v-card-actions>
        <span v-if="question.completed">
          <v-chip
            label
            class="ma-2"
            color="rgb(116, 181, 103)"
            text-color="white"
          >
            解決済みの質問
          </v-chip>
        </span>
        <span v-else>
          <v-chip
            label
            class="ma-2"
            color="rgb(231, 175, 95)"
            text-color="white"
          >
            未解決の質問
          </v-chip>
        </span>
        </v-card-actions>

        <v-card-actions>
          <v-card-text>
            投稿者: {{ userName }}  
            <br>
            投稿日時: {{ question.date }}
            <br>
            <span v-if="question.url">
              URL: <a :href="question.url">{{ question.url }}</a>
            </span>
          </v-card-text>
        </v-btn>
          <v-btn 
            v-if="question.completed"
            :disabled="name != userName"
            @click="updateQuestionCompleted" 
            class="ma-2"
            tile
            outlined 
            color="rgb(223, 177, 109)"
            rounded
          >
            <v-icon>mdi-undo-variant</v-icon>未解決に戻す
          </v-btn>
          <v-btn 
            v-else
            :disabled="name != userName"
            @click="updateQuestionCompleted" 
            class="ma-2"
            tile
            outlined 
            color="rgb(131, 179, 112)"
            rounded
          >
            <v-icon>mdi-check</v-icon>解決済みにする
          </v-btn>
          <v-btn icon 
            color="error"
            :disabled="name != userName"
          >
            <v-icon @click="alert = !alert">mdi-delete</v-icon>
          </v-btn> 
          <v-btn icon
            color="pink"
            :disabled="name === userName || name === ''"
          >
            <v-icon @click="favoriteQuestion">mdi-heart</v-icon>
          </v-btn>
          {{ question.favoriteCount }}
        </v-card-actions>
        <v-row>
          <v-col md="12">
            <v-alert
              outlined
              :value="alert"
              transition="scale-transition"
            >
              <v-col col="12" sm="8">
                <h1>本当に削除しますか?</h1>
                <p>この操作は取り消せません.質問に付与された「いいね」も取り消されます.</p>
              </v-col>
              <v-col col="12" sm="4">
                <v-btn color="error" @click="deleteQuestion">削除する</v-btn>
                &nbsp;
                <v-btn @click="alert = !alert">戻る</v-btn>
              </v-col>
            </v-alert>
          </v-col>
        </v-row>
      </div>
      <div v-else class="text-center">
        <v-progress-circular
          :size="100"
          color="primary"
          indeterminate
        ></v-progress-circular>
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';

@Component({
  components: {
    SquarePanel,
  },
})
export default class QuestionPanelDetail extends Vue {

  // これどうやねん
  private fa: boolean = false;
  private tr: boolean = true;

  // 元々 string として良いのでは ??
  private questionId!: number;

  // データベース Question 通り
  private question: any = {};

  // 質問者の名前
  private userName: string = '';
  // ユーザの名前
  private name: string = '';

  private isReady: boolean = false;

  private alert: boolean = false;

  private created(): void {
    if (this.getToken() != null) {
      const claims = JSON.parse(atob(this.getToken().split('.')[1]));
      this.name = claims.name;
    }
    this.questionId = Number(this.$route.query.questionId);
    this.createQuestion();
    this.browseQuestion();
  }

  private updateQuestionCompleted(): void {
    const url = 'api/question/' + String(this.questionId) + '/completed';
    const method = 'PUT';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      location.reload();
    });
  }

  private favoriteQuestion(): void {
    const url = 'api/question/' + String(this.questionId) + '/favorite';
    const method = 'put';
    const headers = {authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      this.createQuestion();
    });
  }

  private browseQuestion(): void {
    const url = 'api/no-auth/question/' + String(this.questionId) + '/browse';
    const method = 'put';
    fetch(url, {method});

  }

  private createQuestion(): void {
    const url = '/api/no-auth/question/' + String(this.questionId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.question = json;
      this.isReady = true;
      this.setUser();
    });
  }

  private setUser(): void {

    const url = 'api/no-auth/user/' + String(this.question.uid);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
    });
  }

  private deleteQuestion(): void {
    const url = 'api/question/' + String(this.questionId);
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        // 質問を削除しました
        alert('質問を削除しました');
        this.$router.push('./');
      }
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.small {
  font-size: 75%;
}
.question-panel-detail {
}
.mavon-editor {
  height: 100;
  z-index: 0;
}
</style>
